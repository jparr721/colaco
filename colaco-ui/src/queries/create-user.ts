interface UserCreateResponse {
    id: string;
	balance: number;
	is_admin: boolean;
	created_at: string;
	updated_at: string;
}

export default async function createUser(): Promise<UserCreateResponse> {
    const response = await fetch(`http://localhost:8000/v1/users`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({})
    });
    const data = await response.json();
    return data;
}