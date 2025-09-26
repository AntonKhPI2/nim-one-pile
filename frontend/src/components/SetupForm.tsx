import { useState } from 'react'
import type { Variant } from '../api'

export default function SetupForm({ onSubmit }: { onSubmit: (v: Variant, n: number, k: number) => void }){
    const [variant, setVariant] = useState<Variant>('normal')
    const [n, setN] = useState(21)
    const [k, setK] = useState(3)

    return (
        <form onSubmit={e => { e.preventDefault(); onSubmit(variant, n, k) }}>
            <fieldset style={{ display: 'grid', gap: 12 }}>
                <label>
                    Game variant:&nbsp;
                    <select value={variant} onChange={e => setVariant(e.target.value as Variant)}>
                        <option value="normal">Normal (last wins)</option>
                        <option value="misere">Misère (last loses)</option>
                    </select>
                </label>

                <label>
                    Initial coins N:&nbsp;
                    <input
                        type="number"
                        min={1}
                        value={n}
                        onChange={e => setN(parseInt(e.target.value || '1', 10))}
                    />
                </label>

                <label>
                    Max per move k:&nbsp;
                    <input
                        type="number"
                        min={1}
                        value={k}
                        onChange={e => setK(parseInt(e.target.value || '1', 10))}
                    />
                </label>

                <button type="submit">Start</button>
            </fieldset>
        </form>
    )
}
