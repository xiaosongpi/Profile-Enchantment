import { Home, Users, User, Wallet } from "lucide-react"
import SidebarButton from "./SidebarButton"
import CircleButton from "./CircleButton"

function Sidebar() {
    return(
        <div className="min-w-64 h-full border-r border-r-sidebar-fourth/20 bg-linear-to-t from-sidebar-dark from-70% via-[#222325] via-85% to-sidebar-trinary/20 flex flex-col justify-between">
            <div>
                
            </div>
            <div className="p-5 w-full gap-y-2.5 flex flex-col">
                <SidebarButton>
                  <Home className="w-5 h-5" />
                  <span className="text-sm">Home</span>
                </SidebarButton>
                <SidebarButton>
                  <Users className="w-5 h-5" />
                  <span className="text-sm">User</span>
                </SidebarButton>
                <SidebarButton>
                  <Wallet className="w-5 h-5" />
                  <span className="text-sm">Wallet</span>
                </SidebarButton>
            </div>

            <div className="w-full p-5 gap-2.5 border border-transparent border-t-sidebar-fourth/20 flex items-center justify-start">
                <CircleButton>
                    <User className="w-5 h-5" />
                </CircleButton>
                <div className="flex flex-col text-sidebar-fourth/60">
                    <span className="text-sm font-bold">Admin User</span>
                    <span className="text-xs">adminuser@gmail.com</span>
                </div>
                <button className="w-5 h-5 gap-y-0.5 flex flex-col items-center justify-center cursor-pointer">
                    <div className="w-1 h-1 rounded-full bg-sidebar-fourth/20" />
                    <div className="w-1 h-1 rounded-full bg-sidebar-fourth/20" />
                    <div className="w-1 h-1 rounded-full bg-sidebar-fourth/20" />
                </button>
            </div>
        </div>
    )
}

export default Sidebar