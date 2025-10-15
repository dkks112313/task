export interface Event {
    id: number;
    user_id: number;
    action: string;
    metadata: string;
    timestamp: string;
}

export interface EventFilters {
    user_id?: string;
    from?: string;
    to?: string;
    action?: string;
    metadata?: string;
}
