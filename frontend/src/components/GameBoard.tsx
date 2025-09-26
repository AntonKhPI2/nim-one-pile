import { State } from '../api'

const cap = (s: string) => (s ? s[0].toUpperCase() + s.slice(1) : s);

export default function GameBoard({
                                      state,
                                      onTake,
                                  }: {
    state: State
    onTake: (t: number) => void
}) {
    const buttons = Array.from({ length: state.k }, (_, i) => i + 1)

    const isFinished = !!state.winner
    const isHumanTurn = state.player_turn === 'human'
    const disabled = isFinished || !isHumanTurn || state.remaining === 0

    return (
        <div style={{ marginTop: 16 }}>
            <p>
                Coins left: <b>{state.remaining}</b>
            </p>

            {!isFinished ? (
                <p>
                    Turn: <b>{cap(state.player_turn)}</b>
                </p>
            ) : (
                <p>
                    Winner: <b>{cap(state.winner)}</b>
                </p>
            )}

            <div style={{ display: 'flex', flexWrap: 'wrap', gap: 8 }}>
                {buttons.map((t) => (
                    <button
                        key={t}
                        disabled={disabled || t > state.remaining}
                        onClick={() => onTake(t)}
                    >
                        Take {t}
                    </button>
                ))}
            </div>
        </div>
    )
}
