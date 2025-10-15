import type {Event} from "../types/types.ts";

interface EventsTableProps {
    events: Event[];
}

export default function EventsTable({ events }: EventsTableProps) {
    if ( events == null || !events.length) {
        return <p className="p-4 text-gray-600">Don't have events</p>;
    }

    console.log("Received events:", events);

    return (
            <table className="min-w-full border-collapse border mt-4">
                <thead>
                <tr className="bg-gray-100">
                    <th className="border p-2">User ID</th>
                    <th className="border p-2">Action</th>
                    <th className="border p-2">Metadata</th>
                    <th className="border p-2">Time</th>
                </tr>
                </thead>
                <tbody>
                {events.map((e) => (
                    <tr>
                    <td className="border p-2">{e.user_id}</td>
                    <td className="border p-2">{e.action}</td>
                    <td className="border p-2">{typeof e.metadata === "string" ? e.metadata : JSON.stringify(e.metadata)}</td>
                    <td className="border p-2">{new Date(e.timestamp).toLocaleString()}</td>
                    </tr>
                ))}
                </tbody>
            </table>
    );
}
