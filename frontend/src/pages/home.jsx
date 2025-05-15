import Form from "../components/form";
import { generateResume } from "../api/resume";

export default function Home() {
    const handleSubmit = async (data) => {
        try {
            const pdfBlob = await generateResume(data);
            const url = URL.createObjectURL(pdfBlob);
            const link = document.createElement('a');
            link.href = url;
            link.download = 'resume.pdf';
            link.click();
            URL.revokeObjectURL(url);
        } catch (error) {
            alert(error.message);
        }
    };

    return (
        <div>
            <h1 className="text-center text-2x1 font-bold">Resume Generator</h1>
            <Form onSubmit={handleSubmit} />
        </div>
    );
}