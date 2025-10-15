import { useState } from "react";
import Filters from "./components/Filters";
import EventsTable from "./components/EventsTable";
import { fetchEvents } from "./api";
import type {Event, EventFilters} from "./types/types.ts";

export default function App() {
    const [events, setEvents] = useState<Event[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const handleApplyFilters = async (filters: EventFilters) => {
        setLoading(true);
        setError("");
        try {
            const data = await fetchEvents(filters);
            setEvents(data);
        } catch (err) {
            setError("Error when downloading data");
        } finally {
            setLoading(false);
        }
    };

    return (
            <div className="max-w-5xl mx-auto my-8 bg-white shadow rounded">
                <h1 className="text-2xl font-semibold p-4 border-b">
                    View events from user
                </h1>

                <Filters onApply={handleApplyFilters} />

                {loading && <p className="p-4">Download...</p>}
                {error && <p className="p-4 text-red-600">{error}</p>}
                {!loading && !error && <EventsTable events={events} />}
            </div>
    );
}
