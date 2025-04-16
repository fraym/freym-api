export interface PaginatedResponse<T extends object> {
    limit: number;
    page: number;
    total: number;
    data: T[];
}
