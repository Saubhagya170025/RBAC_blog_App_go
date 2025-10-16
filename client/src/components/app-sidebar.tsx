import { UserStar, Home, Settings, ChartBarStacked, ChevronUp, User2 } from "lucide-react"

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarHeader,
  SidebarFooter
} from "@/components/ui/sidebar"

// Optional dropdown imports (only needed if you enable the footer dropdown)
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { redirect } from "react-router-dom";
import axios from "axios";

// Menu items.
const items = [
  { title: "Home", url: "/dashboard", icon: Home },
  { title: "Categories", url: "/dashboard/categories", icon: ChartBarStacked },
  { title: "Roles", url: "/dashboard/roles", icon: UserStar },
  { title: "Settings", url: "#", icon: Settings },
]

const API = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/auth';

export function AppSidebar() {
  async function handleLogout() {
    try {
      const res = await axios.post(`${API}/logout`, {}, {
        headers: { 'Content-Type': 'application/json' },
        withCredentials: true,
      });
      console.log("logout button clicked")
      if (res.status !== 200) {
        throw new Error('Logout failed');
      }
      // Handle successful logout (e.g., redirect to login)
    } catch (error) {
      console.error('Logout error:', error);
    }
    redirect('/home');
  }
  return (
    <Sidebar>
      <SidebarHeader>
        <h1 className="font-bold">Logo here</h1>
      </SidebarHeader>

      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Application</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <a href={item.url} className="flex items-center gap-3">
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      {/* Optional footer with dropdown - enable if dropdown-menu component exists */}

      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <SidebarMenuButton>
                  <User2 /> Username
                  <ChevronUp className="ml-auto" />
                </SidebarMenuButton>
              </DropdownMenuTrigger>
              <DropdownMenuContent side="top" className="w-[--radix-popper-anchor-width]">
                <DropdownMenuItem>
                  <span>Account</span>
                </DropdownMenuItem>
                <DropdownMenuItem>
                  <span>Billing</span>
                </DropdownMenuItem>
                <DropdownMenuItem>
                  <span onClick={handleLogout}>Sign out</span>
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>

    </Sidebar>
  )
}
