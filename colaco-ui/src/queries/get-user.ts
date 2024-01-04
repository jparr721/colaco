
export interface User {
	id:        string;
	balance:   number;
	isAdmin:   boolean;
	createdAt: string;
	updatedAt: string;
}

export default async function getUser({ queryKey }: { queryKey: [string, string] }): Promise<User> {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [_, id] = queryKey;
    const response = await fetch(
        `http://localhost:8000/v1/users/me`,
        {
            headers: {
                'x-auth-token': id
            }
        }
    );
    const data = await response.json();
    return data;
}