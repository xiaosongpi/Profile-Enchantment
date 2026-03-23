import { Users, UserRoundCheck, Search, PackageOpenIcon, Plus } from "lucide-react"

function User() {
    return (
        <div className="p-6 w-full flex flex-col gap-5">
            <div className="flex justify-between items-center">
                <div className="flex flex-col gap-0.5 pb-3">
                    <h1 className="text-sidebar-fourth/60 text-2xl font-bold">User Panel</h1>
                    <p className="text-sidebar-fourth/40">Managed and monitoring user</p>
                </div>
                <button className="h-10 px-5 py-2.5 gap-2.5 flex items-center justify-center border border-transparent bg-sidebar-primary shadow-[3px_3px_6px_#000000] rounded-2xl cursor-pointer transition duration-500 active:shadow active:bg-sidebar-dark active:border active:border-sidebar-fourth/20">
                    <Plus className="w-5 h-5 text-sidebar-fourth/60"></Plus>
                    <p className="text-sidebar-fourth/60">Add User</p>
                </button>
            </div>
            <div className="flex gap-5">
                <div className="h-full min-w-40 w-full p-5 gap-3.5 flex bg-sidebar-primary rounded-2xl shadow-[3px_3px_6px_#000000]">
                    <div className="h-14 min-w-14 bg-sidebar-dark rounded-xl flex items-center justify-center shadow-[inset_3px_3px_6px_#000000]">
                        <Users className="w-5 h-5 text-sidebar-fourth/60"></Users>
                    </div>
                    <div className="flex flex-col">
                        <p className="text-sm text-sidebar-fourth/60 overflow-hidden whitespace-nowrap text-ellipsis">All User</p>
                        <p className="text-sidebar-fourth text-2xl font-bold overflow-hidden whitespace-nowrap text-ellipsis">100</p>
                    </div>
                </div>
                <div className="h-full min-w-40 w-full p-5 gap-3.5 flex bg-sidebar-primary rounded-2xl shadow-[3px_3px_6px_#000000]">
                    <div className="h-14 min-w-14 bg-sidebar-dark rounded-xl flex items-center justify-center shadow-[inset_3px_3px_6px_#000000]">
                        <UserRoundCheck className="w-5 h-5 text-sidebar-fourth/60"></UserRoundCheck>
                    </div>
                    <div className="flex flex-col min-w-0">
                        <p className="text-sm w-auto text-sidebar-fourth/60 overflow-hidden whitespace-nowrap text-ellipsis">Today Logged In User</p>
                        <p className="text-sidebar-fourth text-2xl font-bold overflow-hidden whitespace-nowrap text-ellipsis">23</p>
                    </div>
                </div>
                <div className="h-full min-w-40 w-full p-5 gap-3.5 flex bg-sidebar-primary rounded-2xl shadow-[3px_3px_6px_#000000]">
                    <div className="h-14 min-w-14 bg-sidebar-dark rounded-xl flex items-center justify-center shadow-[inset_3px_3px_6px_#000000]">
                        <UserRoundCheck className="w-5 h-5 text-sidebar-fourth/60"></UserRoundCheck>
                    </div>
                    <div className="flex flex-col min-w-0">
                        <p className="text-sm w-auto text-sidebar-fourth/60 overflow-hidden whitespace-nowrap text-ellipsis">Current Actived User</p>
                        <p className="text-sidebar-fourth text-2xl font-bold overflow-hidden whitespace-nowrap text-ellipsis">10</p>
                    </div>
                </div>
            </div>
            <div className="h-full w-full p-5 bg-sidebar-primary rounded-2xl shadow-[3px_3px_6px_#000000]">
                <div className="h-10 px-3 w-full mb-5 bg-sidebar-dark rounded-2xl flex items-center shadow-[inset_3px_3px_6px_#000000]">
                    <button className="h-full mr-2.5">
                        <Search className="w-5 h-5 text-sidebar-fourth/60"></Search>
                    </button>
                    <input 
                        className="h-10 w-full text-sidebar-fourth/60 bg-transparent outline-none border-none placeholder:text-sidebar-fourth/20" 
                        type="text" 
                        placeholder="Search something here..." 
                    />
                     </div>
                <table className="w-full">
                    <thead>
                        <tr className="border-b border-sidebar-fourth/20">
                            <th className="font-medium text-left text-sidebar-fourth/60 text-sm py-3 px-4">Name</th>
                            <th className="font-medium text-left text-sidebar-fourth/60 text-sm py-3 px-4">Email</th>
                            <th className="font-medium text-left text-sidebar-fourth/60 text-sm py-3 px-4">Role</th>
                            <th className="font-medium text-left text-sidebar-fourth/60 text-sm py-3 px-4">Status</th>
                            <th className="font-medium text-left text-sidebar-fourth/60 text-sm py-3 px-4">Last Login</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td colSpan={5} className="py-10">
                                <div className="w-full flex items-center justify-center flex-col gap-2.5">
                                    <PackageOpenIcon className="w-10 h-10 text-sidebar-fourth/60" strokeWidth={1}></PackageOpenIcon>
                                    <p className="text-sm text-sidebar-fourth/60">No User Found</p>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    )
}

export default User