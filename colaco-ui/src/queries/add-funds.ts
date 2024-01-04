export default async function addFunds(id: string, amount: number): Promise<void> {
    await fetch(`http://localhost:8000/v1/users/me/deposit`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'x-auth-token': id,
        },
        body: JSON.stringify({
            amount: amount,
        })
    });
}