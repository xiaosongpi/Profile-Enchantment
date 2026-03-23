import Sidebar from "./sidebar";
import { Outlet } from "react-router-dom"

function Layout() {
    return (
        <div className="bg-sidebar-dark">
            <div className="h-screen flex gap-4">
                <Sidebar />
                <main className="flex-1 overflow-hidden rounded-4xl">
                    <Outlet />
                </main>
            </div>
        </div>
    )
}

export default Layout