import { useState } from 'react'
import { newGame, takeMove, getGame } from './api'
import type { State, Variant } from './api'
import SetupForm from './components/SetupForm'
import GameBoard from './components/GameBoard'
import StatusBar from './components/StatusBar'

export default function App(){
    const [state, setState] = useState<State | null>(null)
    const [error, setError] = useState<string>('')

    const handleSetup = async (variant: Variant, n: number, k: number) => {
        setError('')
        try {
            const s = await newGame({ variant, n, k })
            setState(s)
        } catch (e:any) { setError(e.message) }
    }

    const handleTake = async (take: number) => {
        if (!state) return
        setError('')
        try {
            const s = await takeMove({ id: state.id, take });
            setState(s)
        } catch (e:any) { setError(e.message) }
    }

    return (
        <div style={{ maxWidth: 720, margin: '24px auto', fontFamily: 'system-ui, sans-serif' }}>
            <h1>Nim: one pile</h1>
            {!state && <SetupForm onSubmit={handleSetup} />}
            {state && (
                <>
                    <StatusBar state={state} onRestart={() => setState(null)} />
                    <GameBoard state={state} onTake={handleTake} />
                </>
            )}
            {error && <p style={{ color: 'crimson' }}>{error}</p>}
            <hr />
            <p style={{ opacity: 0.8 }}>
                Strategy: normal — leave a multiple of (k+1); misère — leave 1 (mod k+1) and avoid taking all coins.
            </p>
        </div>
    )
}
