function SidebarButton({ children, onClick }) {
    return (
        <button 
            className="w-full py-2.5 px-5 gap-3 border border-transparent rounded-full flex items-center justify-start transition duration-500 shadow-none cursor-pointer text-sidebar-fourth/60 hover:border hover:border-sidebar-fourth/20 hover:shadow-[3px_3px_6px_#000000]"
            onClick={onClick}
        >
            {children}
        </button>
    )
}

export default SidebarButton