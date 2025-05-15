export async function generateResume(formData) {
    const response = await fetch('/api/generate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData),
    });
    if (!response.ok) {
        throw new Error('Failed to generate resume');
    }
    return response.blob(); // assuming PDF
}