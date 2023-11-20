
export interface ErrorResponse {
    code?: string;
    message?: string;
}

export interface RestApiResponse<T> {
    data?: T;
    error?: ErrorResponse;
}
