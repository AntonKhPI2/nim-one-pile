export type Variant = 'normal' | 'misere';
export type Side = 'human' | 'computer';

export type State = {
    id: string;
    variant: Variant;
    k: number;
    remaining: number;
    player_turn: Side;
    winner: '' | Side;
};

export type NewGameReq = {
    variant: Variant;
    n: number;
    k: number;
};

export type TakeReq = {
    id: string;
    take: number;
};

const BASE = (import.meta.env.VITE_API_BASE_URL ?? '/api').replace(/\/+$/, '');

async function json<T>(path: string, init?: RequestInit): Promise<T> {
    const res = await fetch(`${BASE}${path}`, {
        headers: { 'Content-Type': 'application/json', ...(init?.headers || {}) },
        ...init,
    });
    if (!res.ok) {
        const text = await res.text().catch(() => res.statusText);
        throw new Error(text || `HTTP ${res.status}`);
    }
    return res.json() as Promise<T>;
}

export function newGame(payload: NewGameReq) {
    return json<State>('/new-game', {
        method: 'POST',
        body: JSON.stringify(payload),
    });
}

export function getGame(id: string) {
    return json<State>(`/game/${encodeURIComponent(id)}`);
}

export function takeMove(payload: TakeReq) {
    return json<State>('/take', {
        method: 'POST',
        body: JSON.stringify(payload),
    });
}
