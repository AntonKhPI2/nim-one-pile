import { State } from '../api';

const cap = (s: string) => (s ? s[0].toUpperCase() + s.slice(1) : s);

export default function StatusBar({ state, onRestart }: { state: State; onRestart: () => void }) {
    const variantLabel = state.variant === 'misere' ? 'Misère' : 'Normal';

    return (
        <div style={{ background: '#f4f4f4', padding: 16, borderRadius: 8, marginBottom: 16 }}>
            <div><strong>Mode:</strong> {variantLabel} | k = {state.k}</div>

            {state.winner && (
                <div style={{ marginTop: 8 }}>
                    <strong>Winner:</strong> {cap(state.winner)}
                    <button onClick={onRestart} style={{ marginLeft: 8 }}>New game</button>
                </div>
            )}
        </div>
    );
}
