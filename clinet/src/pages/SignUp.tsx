import { useEffect, useRef, useState } from "react"
import { FaSpinner } from 'react-icons/fa';
import { useNavigate } from "react-router-dom";

// Define structure of form data
interface SignUpFormData {
    first_name: string
    last_name: string
    email: string
    username: string
    password: string
}

export default function SignUp() {
    const navigate = useNavigate()

    // Set initial value of form data
    const [formData, setFormData] = useState<SignUpFormData>({
        first_name: '',
        last_name: '',
        email: '',
        username: '',
        password: '',
    })

    // States to manage UI
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [message, setMessage] = useState<string | null>(null)
    const [isError, setIsError] = useState<boolean>(false)

    // Use a 'ref' to focus the first input on load
    const firstnameInputRef = useRef<HTMLInputElement>(null)

    useEffect(() => {
        if (firstnameInputRef.current) {
            firstnameInputRef.current.focus()
        }
    }, [])

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target

        setFormData({
            ...formData,
            [name]: value,
        })
    }
    
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        setIsSubmitting(true)
        setMessage(null)
        setIsError(false)

        try {
            const res = await fetch('/api/auth/signup', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(formData)
            })

            const data = await res.json()
            
            if (data.status === 'error') {
                setIsError(true)
                // Use the 'error' field sent from your backend
                setMessage(data.error || "An error occurred during signup.")
                setIsSubmitting(false)
                return;
            }
            
            navigate("/")
        }
        catch (error) {
            setIsError(true)
            setMessage(error instanceof Error ? error.message : "An unexpected error occurred")
            setIsSubmitting(false)
        }
    }

    const inputClasses = "w-full px-4 py-2.5 text-gray-900 bg-gray-50 border border-gray-200 rounded-lg outline-none transition-all focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 focus:bg-white"

    return (
        <div className="flex items-center justify-center min-h-screen px-4 bg-gray-50">
            <div className="w-full max-w-md p-8 space-y-6 bg-white border border-gray-100 shadow-xl rounded-2xl">
                <div className="space-y-2 text-center">
                    <h2 className="text-3xl font-bold tracking-tight text-gray-900">Create an account</h2>
                    <p className="text-sm text-gray-500">Please fill in your details to sign up.</p>
                </div>
        
                <form onSubmit={handleSubmit} className="space-y-5">
                    {/* First & Last Name Row */}
                    <div className="grid grid-cols-1 gap-5 sm:grid-cols-2">
                        <div className="space-y-1.5">
                            <label htmlFor="first_name" className="block text-sm font-medium text-gray-700">First Name</label>
                            <input
                                type="text"
                                id="first_name"
                                name="first_name"
                                ref={firstnameInputRef}
                                value={formData.first_name}
                                onChange={handleChange}
                                required
                                className={inputClasses}
                                placeholder="John"
                            />
                        </div>
                        <div className="space-y-1.5">
                            <label htmlFor="last_name" className="block text-sm font-medium text-gray-700">Last Name</label>
                            <input
                                type="text"
                                id="last_name"
                                name="last_name"
                                value={formData.last_name}
                                onChange={handleChange}
                                required
                                className={inputClasses}
                                placeholder="Doe"
                            />
                        </div>
                    </div>
        
                    {/* Username */}
                    <div className="space-y-1.5">
                        <label htmlFor="username" className="block text-sm font-medium text-gray-700">Username</label>
                        <input
                            type="text"
                            id="username"
                            name="username"
                            value={formData.username}
                            onChange={handleChange}
                            required
                            className={inputClasses}
                            placeholder="johndoe123"
                        />
                    </div>
        
                    {/* Email */}
                    <div className="space-y-1.5">
                        <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email Address</label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            required
                            className={inputClasses}
                            placeholder="john@example.com"
                        />
                    </div>
        
                    {/* Password */}
                    <div className="space-y-1.5">
                        <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            value={formData.password}
                            onChange={handleChange}
                            required
                            className={inputClasses}
                            placeholder="••••••••"
                        />
                    </div>
        
                    {/* Submit Button */}
                    <button
                        type="submit"
                        disabled={isSubmitting}
                        className="flex items-center justify-center w-full px-4 py-3 mt-4 text-sm font-semibold text-white transition-colors bg-indigo-600 rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-4 focus:ring-indigo-200 disabled:opacity-70 disabled:cursor-not-allowed"
                    >
                        {isSubmitting ? (
                            <>
                                <FaSpinner className="w-5 h-5 mr-3 -ml-1 animate-spin" />
                                Creating account...
                            </>
                        ) : (
                            'Sign Up'
                        )}
                    </button>

                    {/* Alert Message Placed Below Button */}
                    {message && (
                        <div 
                            className={`p-3 mt-4 text-sm text-center rounded-lg border ${
                                isError 
                                    ? 'text-red-800 bg-red-50 border-red-200' 
                                    : 'text-green-800 bg-green-50 border-green-200'
                            }`}
                            role="alert"
                        >
                            {message}
                        </div>
                    )}
                </form>
            </div>
        </div>
    )
}
