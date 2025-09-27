import {fireEvent, render, screen, waitFor} from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import App from '../App'
import * as api from '../api'

vi.mock('../api')

const mockedApi = vi.mocked(api)

const mockInitialState: api.State = {
    id: 'game123',
    variant: 'normal',
    k: 3,
    remaining: 21,
    player_turn: 'human',
    winner: '',
}

describe('App', () => {
    beforeEach(() => {
        vi.resetAllMocks()
    })

    it('should render the setup form initially', () => {
        render(<App />)
        expect(screen.getByRole('button', { name: /Start/i })).toBeInTheDocument()
        expect(screen.queryByText(/Coins left:/i)).not.toBeInTheDocument()
    })

    it('should start a new game and render the game board', async () => {
        const user = userEvent.setup()
        mockedApi.newGame.mockResolvedValue(mockInitialState)
        render(<App />)

        const nInput = screen.getByLabelText(/Initial coins N/i);
        fireEvent.change(nInput, { target: { value: '21' } })

        await user.click(screen.getByRole('button', { name: /Start/i }))

        expect(mockedApi.newGame).toHaveBeenCalledWith({ variant: 'normal', n: 21, k: 3 })


        await waitFor(() => {
            expect(screen.getByText(/Coins left:/)).toHaveTextContent('Coins left: 21')
        })

        expect(screen.queryByRole('button', { name: /Start/i })).not.toBeInTheDocument()
    })

    it('should handle a player move', async () => {
        const user = userEvent.setup()
        mockedApi.newGame.mockResolvedValue(mockInitialState)

        const stateAfterMove: api.State = { ...mockInitialState, remaining: 18, player_turn: 'computer' }
        mockedApi.takeMove.mockResolvedValue(stateAfterMove)

        render(<App />)

        await user.click(screen.getByRole('button', { name: /Start/i }))
        await screen.findByText(/Coins left:/)
        expect(screen.getByText(/Coins left:/)).toHaveTextContent("21")

        const takeButton = screen.getByRole('button', { name: /Take 3/i })
        await user.click(takeButton)

        expect(mockedApi.takeMove).toHaveBeenCalledWith({ id: 'game123', take: 3 })

        await waitFor(() => {
            expect(screen.getByText(/Coins left:/)).toHaveTextContent('Coins left: 18')
            expect(screen.getByText((content, element) =>
                element?.tagName.toLowerCase() === 'p' && content.startsWith('Turn:')
            )).toHaveTextContent('Turn: Computer')
        })
    })

    it('should handle game end and restart', async () => {
        const user = userEvent.setup()
        mockedApi.newGame.mockResolvedValue(mockInitialState)

        const winningState: api.State = { ...mockInitialState, remaining: 0, player_turn: 'human', winner: 'computer' }
        mockedApi.takeMove.mockResolvedValue(winningState)

        render(<App />)

        await user.click(screen.getByRole('button', { name: /Start/i }))
        await screen.findByText(/Coins left:/)
        expect(screen.getByText(/Coins left:/)).toHaveTextContent("21")
        await user.click(screen.getByRole('button', { name: /Take 1/i }))

        await waitFor(() => {
            expect(screen.getByText((content, element) =>
                element?.tagName.toLowerCase() === 'p' && content.startsWith('Winner:')
            )).toHaveTextContent('Winner: Computer')
        })

        const restartButton = screen.getByRole('button', { name: /New game/i })
        await user.click(restartButton)

        expect(screen.getByRole('button', { name: /Start/i })).toBeInTheDocument()
    })

    it('should display an error message if API call fails', async () => {
        const user = userEvent.setup()
        mockedApi.newGame.mockRejectedValue(new Error('Network Error'))

        render(<App />)
        await user.click(screen.getByRole('button', { name: /Start/i }))
        await waitFor(() => {
            expect(screen.getByText('Network Error')).toBeInTheDocument()
        })
    })
})