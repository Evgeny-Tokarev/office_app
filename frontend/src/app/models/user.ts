export interface User {
    id: number,
    photo?: string,
    name: string,
    role: string,
    password: string,
    email: string,
    registered_at?: string,
}
