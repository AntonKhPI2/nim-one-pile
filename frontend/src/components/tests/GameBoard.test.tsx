import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, it, expect, vi } from 'vitest'
import GameBoard from '../GameBoard'
import type { State } from '../../api'

const mockState: State = {
    id: 'test-game',
    variant: 'normal',
    k: 3,
    remaining: 10,
    player_turn: 'human',
    winner: '',
}

describe('GameBoard', () => {
    it('renders current game state', () => {
        render(<GameBoard state={mockState} onTake={() => {}} />)

        expect(screen.getByText((content, element) => {
            return element?.tagName.toLowerCase() === 'p' && content.startsWith('Coins left:')
        })).toHaveTextContent('Coins left: 10')

        expect(screen.getByText((content, element) => {
            return element?.tagName.toLowerCase() === 'p' && content.startsWith('Turn:')
        })).toHaveTextContent('Turn: Human')

    })

    it('renders winner instead of turn when game is finished', () => {
        const finishedState = { ...mockState, winner: 'computer' as const }
        render(<GameBoard state={finishedState} onTake={() => {}} />)

        expect(screen.getByText((content, element) => {
            return element?.tagName.toLowerCase() === 'p' && content.startsWith('Winner:')
        })).toHaveTextContent('Winner: Computer')

        expect(screen.queryByText(/Turn:/i)).not.toBeInTheDocument()

    })

    it('enables buttons when it is human turn', () => {
        render(<GameBoard state={mockState} onTake={() => {}} />)

        expect(screen.getByRole('button', { name: 'Take 1' })).toBeEnabled()
        expect(screen.getByRole('button', { name: 'Take 2' })).toBeEnabled()
        expect(screen.getByRole('button', { name: 'Take 3' })).toBeEnabled()
    })

    it('disables buttons when it is not human turn', () => {
        const computerTurnState = { ...mockState, player_turn: 'computer' as const }
        render(<GameBoard state={computerTurnState} onTake={() => {}} />)

        expect(screen.getByRole('button', { name: 'Take 1' })).toBeDisabled()
        expect(screen.getByRole('button', { name: 'Take 2' })).toBeDisabled()
        expect(screen.getByRole('button', { name: 'Take 3' })).toBeDisabled()
    })

    it('disables buttons for moves that are not possible', () => {
        const lowRemainingState = { ...mockState, remaining: 2 }
        render(<GameBoard state={lowRemainingState} onTake={() => {}} />)

        expect(screen.getByRole('button', { name: 'Take 1' })).toBeEnabled()
        expect(screen.getByRole('button', { name: 'Take 2' })).toBeEnabled()
        expect(screen.getByRole('button', { name: 'Take 3' })).toBeDisabled()
    })

    it('calls onTake with the correct number when a button is clicked', async () => {
        const user = userEvent.setup()
        const handleTake = vi.fn()
        render(<GameBoard state={mockState} onTake={handleTake} />)

        await user.click(screen.getByRole('button', { name: 'Take 2' }))

        expect(handleTake).toHaveBeenCalledTimes(1)
        expect(handleTake).toHaveBeenCalledWith(2)
    })
})