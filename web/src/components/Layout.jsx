import Sidebar from "./sidebar";
import { Outlet } from "react-router-dom"

function Layout() {
    return (
        <body className="bg-sidebar-primary">
            <div className="p-4 h-screen flex gap-4 overflow-hidden">
                <Sidebar />
                <main className="w-screen rounded-4xl">
                    <Outlet />
                </main>
            </div>
        </body>
    )
}

export default Layout