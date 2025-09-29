import { useState } from 'react'
import { newGame, takeMove, getGame } from './api'
import type { State, Variant } from './api'
import SetupForm from './components/SetupForm'
import GameBoard from './components/GameBoard'
import StatusBar from './components/StatusBar'

type LogEntry = { who: 'Computer' | 'Human'; delta: number }

export default function App(){
    const [state, setState] = useState<State | null>(null)
    const [error, setError] = useState<string>('')
    const [logs, setLogs] = useState<LogEntry[]>([])

    const handleSetup = async (variant: Variant, n: number, k: number) => {
        setError('')
        setLogs([]) 
        try {
            const s = await newGame({ variant, n, k })

            
            const tookByComputer = Math.max(0, n - s.remaining)
            if (tookByComputer > 0) {
                setLogs([{ who: 'Computer', delta: -tookByComputer }])
            }

            setState(s)
        } catch (e:any) { setError(e.message) }
    }

    const handleTake = async (take: number) => {
        if (!state) return
        setError('')
        const prevRemaining = state.remaining

        try {
            const s = await takeMove({ id: state.id, take });
            const newLogs: LogEntry[] = [{ who: 'Human', delta: -take }]
            const compDelta = (prevRemaining - take) - s.remaining
            if (compDelta > 0) {
                newLogs.push({ who: 'Computer', delta: -compDelta })
            }
            setLogs(curr => [...curr, ...newLogs])
            setState(s)
        } catch (e:any) { setError(e.message) }
    }

    const handleRestart = () => {
        setState(null)
        setLogs([])
    }

    return (
        <div style={{ maxWidth: 720, margin: '24px auto', fontFamily: 'system-ui, sans-serif' }}>
            <h1>Nim: one pile</h1>
            {!state && <SetupForm onSubmit={handleSetup} />}
            {state && (
                <>
                    <StatusBar state={state} onRestart={handleRestart} />
                    <GameBoard state={state} onTake={handleTake} />
                </>
            )}
            {error && <p style={{ color: 'crimson' }}>{error}</p>}

            {}
            {logs.length > 0 && (
                <div style={{ marginTop: 24 }}>
                    <h3>Game history</h3>
                    <ul style={{ listStyle: 'none', padding: 0, margin: 0, display: 'grid', gap: 6 }}>
                        {logs.map((e, i) => (
                            <li key={i} style={{ fontFamily: 'ui-monospace, SFMono-Regular, Menlo, monospace' }}>
                                {e.who}: {e.delta}
                            </li>
                        ))}
                    </ul>
                </div>
            )}
            <hr />
            <p style={{ opacity: 0.8 }}>
                Strategy: normal — leave a multiple of (k+1); misère — leave 1 (mod k+1) and avoid taking all coins.
            </p>
        </div>
    )
}
