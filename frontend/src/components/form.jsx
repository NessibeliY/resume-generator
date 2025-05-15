import { useState } from 'react';

export default function Form({ onSubmit }) {
    const [formData, setFormData] = useState({
        name: '',
        experience: '',
        skills: '',
    });

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(formData);
    };

    return (
        <form onSubmit={handleSubmit} className="flex flex-col gap-4 max-w-md mx-auto">
            <input name="name" placeholder="Name" value={formData.name} onChange={handleChange}/>
            <textarea name="experience" placeholder="Experience" value={formData.experience} onChange={handleChange}/>
            <textarea name="skills" placeholder="Skills" value={formData.skills} onChange={handleChange}/>
            <button type="submit">Generate Resume</button>
        </form>
    );
}