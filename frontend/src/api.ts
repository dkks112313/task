import type {Event, EventFilters} from "./types/types.ts";

export async function fetchEvents(filters: EventFilters): Promise<Event[]> {
    const params = new URLSearchParams(
            Object.entries(filters).filter(([_, v]) => v && v !== "")
    );

    const res = await fetch(`http://localhost:8080/events?${params.toString()}`);

    if (!res.ok) {
        throw new Error(`Error when give events (${res.status})`);
    }

    return res.json();
}
