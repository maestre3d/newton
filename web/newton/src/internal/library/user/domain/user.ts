// User Newton user 
export interface User {
    ID: string
    DisplayName: string
    Username: string
    PreferredUsername: string
    Image?: string
}

export const DefaultUser: User = {
    ID: 'a18e69a5-0043-4806-9497-b875ef20b8df',
    DisplayName: 'Alonso Ruiz',
    Username: 'aruizmx',
    PreferredUsername: 'aruizea',
}