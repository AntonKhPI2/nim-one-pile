import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, it, expect, vi } from 'vitest'
import StatusBar from '../StatusBar'
import type { State } from '../../api'

const mockState: State = {
    id: 'test-game',
    variant: 'normal',
    k: 3,
    remaining: 10,
    player_turn: 'human',
    winner: '',
}

describe('StatusBar', () => {
    it('renders game mode and k value during the game', () => {
        render(<StatusBar state={mockState} onRestart={() => {}} />)

        expect(screen.getByText(/Mode: Normal | k = 3/i)).toBeInTheDocument()
    })

    it('displays "Misère" for misere variant', () => {
        const misereState = { ...mockState, variant: 'misere' as const }
        render(<StatusBar state={misereState} onRestart={() => {}} />)

        expect(screen.getByText(/Mode: Misère | k = 3/i)).toBeInTheDocument()
    })

    it('does not show winner info or restart button when game is ongoing', () => {
        render(<StatusBar state={mockState} onRestart={() => {}} />)

        expect(screen.queryByText(/Winner:/i)).not.toBeInTheDocument()
        expect(screen.queryByRole('button', { name: /New game/i })).not.toBeInTheDocument()
    })

    it('shows winner info and restart button when game is finished', () => {
        const finishedState = { ...mockState, winner: 'human' as const }
        render(<StatusBar state={finishedState} onRestart={() => {}} />)

        const winnerDisplay = screen.getByText(/Winner:/i).parentElement;
        expect(winnerDisplay).toHaveTextContent('Winner: Human');

        expect(screen.getByRole('button', { name: /New game/i })).toBeInTheDocument()

    })

    it('calls onRestart when "New game" button is clicked', async () => {
        const user = userEvent.setup()
        const handleRestart = vi.fn()
        const finishedState = { ...mockState, winner: 'computer' as const }
        render(<StatusBar state={finishedState} onRestart={handleRestart} />)

        const restartButton = screen.getByRole('button', { name: /New game/i })
        await user.click(restartButton)

        expect(handleRestart).toHaveBeenCalledTimes(1)
    })
})