import { render, screen, fireEvent } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, it, expect, vi } from 'vitest'
import SetupForm from '../SetupForm'

describe('SetupForm', () => {
    it('renders initial form fields correctly', () => {
        render(<SetupForm onSubmit={() => {}} />)

        expect(screen.getByLabelText(/Game variant/i)).toHaveValue('normal')
        expect(screen.getByLabelText(/Initial coins N/i)).toHaveValue(21)
        expect(screen.getByLabelText(/Max per move k/i)).toHaveValue(3)
        expect(screen.getByRole('button', { name: /Start/i })).toBeInTheDocument()
    })

    it('updates form fields on user input', async () => {
        const user = userEvent.setup()
        render(<SetupForm onSubmit={() => {}} />)

        const variantSelect = screen.getByLabelText(/Game variant/i)
        await user.selectOptions(variantSelect, 'misere')
        expect(variantSelect).toHaveValue('misere')

        const nInput = screen.getByLabelText(/Initial coins N/i)
        fireEvent.change(nInput, { target: { value: '50' } })
        expect(nInput).toHaveValue(50)

        const kInput = screen.getByLabelText(/Max per move k/i)
        fireEvent.change(kInput, { target: { value: '5' } })
        expect(kInput).toHaveValue(5)
    })

    it('calls onSubmit with form values when submitted', async () => {
        const user = userEvent.setup()
        const handleSubmit = vi.fn()
        render(<SetupForm onSubmit={handleSubmit} />)

        await user.selectOptions(screen.getByLabelText(/Game variant/i), 'misere')

        const nInput = screen.getByLabelText(/Initial coins N/i)
        fireEvent.change(nInput, { target: { value: '30' } })

        const kInput = screen.getByLabelText(/Max per move k/i)
        fireEvent.change(kInput, { target: { value: '4' } })

        await user.click(screen.getByRole('button', { name: /Start/i }))

        expect(handleSubmit).toHaveBeenCalledTimes(1)
        expect(handleSubmit).toHaveBeenCalledWith('misere', 30, 4)
    })
})