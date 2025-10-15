import type {Event} from "../types/types.ts";

interface EventsTableProps {
    events: Event[];
}

export default function EventsTable({ events }: EventsTableProps) {
    if (!events.length) {
        return <p className="p-4 text-gray-600">Don't have events</p>;
    }

    return (
            <table className="min-w-full border-collapse border mt-4">
                <thead>
                <tr className="bg-gray-100">
                    <th className="border p-2">ID</th>
                    <th className="border p-2">User ID</th>
                    <th className="border p-2">Action</th>
                    <th className="border p-2">Metadata</th>
                    <th className="border p-2">Time</th>
                </tr>
                </thead>
                <tbody>
                {events.map((e) => (
                        <tr key={e.id}>
                            <td className="border p-2">{e.id}</td>
                            <td className="border p-2">{e.user_id}</td>
                            <td className="border p-2">{e.action}</td>
                            <td className="border p-2">{e.metadata}</td>
                            <td className="border p-2">
                                {new Date(e.time_event).toLocaleString()}
                            </td>
                        </tr>
                ))}
                </tbody>
            </table>
    );
}
