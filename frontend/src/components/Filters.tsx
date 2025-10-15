import { useState } from "react";
import type {EventFilters} from "../types/types.ts";

interface FiltersProps {
    onApply: (filters: EventFilters) => void;
}

export default function Filters({ onApply }: FiltersProps) {
    const [filters, setFilters] = useState<EventFilters>({
        user_id: "",
        from: "",
        to: "",
        action: "",
        metadata: "",
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFilters({ ...filters, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        onApply(filters);
    };

    return (
            <form
                    onSubmit={handleSubmit}
                    className="flex flex-wrap gap-3 items-end p-4 border-b"
            >
                <div>
                    <label className="block text-sm">User ID</label>
                    <input
                            name="user_id"
                            type="text"
                            value={filters.user_id}
                            onChange={handleChange}
                            className="border rounded px-2 py-1"
                    />
                </div>

                <div>
                    <label className="block text-sm">From</label>
                    <input
                            name="from"
                            type="date"
                            value={filters.from}
                            onChange={handleChange}
                            className="border rounded px-2 py-1"
                    />
                </div>

                <div>
                    <label className="block text-sm">To</label>
                    <input
                            name="to"
                            type="date"
                            value={filters.to}
                            onChange={handleChange}
                            className="border rounded px-2 py-1"
                    />
                </div>

                <div>
                    <label className="block text-sm">Action</label>
                    <input
                            name="action"
                            type="text"
                            value={filters.action}
                            onChange={handleChange}
                            className="border rounded px-2 py-1"
                    />
                </div>

                <div>
                    <label className="block text-sm">Metadata</label>
                    <input
                            name="metadata"
                            type="text"
                            value={filters.metadata}
                            onChange={handleChange}
                            className="border rounded px-2 py-1"
                    />
                </div>

                <button
                        type="submit"
                        className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700"
                >
                    Use
                </button>
            </form>
    );
}
