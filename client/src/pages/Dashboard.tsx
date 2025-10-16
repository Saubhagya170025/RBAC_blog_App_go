import { SidebarProvider, SidebarTrigger } from '../components/ui/sidebar';
import { AppSidebar } from '../components/app-sidebar';

function Dashboard(){
        return (
                <SidebarProvider >
                        <AppSidebar />
                        <SidebarTrigger />
                        <div className="ml-64 p-6">
                                <h1 className="text-2xl font-bold">Dashboard</h1>
                                <p className="mt-4">Welcome to your dashboard.</p>
                        </div>
                </SidebarProvider>
        )
}

export default Dashboard;
