// Assuming you have an image in your src/assets folder
import { useNavigate } from 'react-router-dom';

function Home() {

    const navigate = useNavigate();
    
       const handleLoginClick = () => {
        // This function navigates the user to the '/login' path
        navigate('/login');
    };

    return (
        <div className="min-h-screen bg-gray-50 font-sans">
            {/* --- Fully Responsive Professional Navbar --- */}
            <nav className="bg-white shadow-md">
                <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div className="flex justify-between items-center h-16">
                        
                        {/* Logo/Brand Name */}
                        <div className="flex-shrink-0">
                            <a href="/" className="text-2xl font-bold text-indigo-600">
                                 Logo
                            </a>
                        </div>

                        {/* Desktop Navigation Links and Login Button (Hidden on Mobile) */}
                        <div className="hidden sm:flex sm:items-center sm:space-x-8">
                            <button 
                                type="button"
                                onClick={handleLoginClick}
                                className="ml-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 transition duration-150"
                            >
                                Login
                            </button>
                        </div>
                        
                      
                    </div>
                </div>
            </nav>
            <header className="relative bg-gray-800 h-[60vh] overflow-hidden">
                <div 
                    className="absolute inset-0 bg-cover bg-center opacity-70" 
                    style={{ backgroundImage: "url('https://plus.unsplash.com/premium_photo-1720744786849-a7412d24ffbf?ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&q=80&w=1109')" }}
                    aria-hidden="true"
                ></div>
            </header>
        </div>
    );
}


export default Home;