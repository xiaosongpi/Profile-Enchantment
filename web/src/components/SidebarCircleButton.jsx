function SidebarCircleButton({children, onClick}) {
    return (
        <button 
            className="p-2.5 rounded-full flex items-center justify-center text-sidebar-fourth/60 transition duration-500 border border-sidebar-fourth/20 shadow-[3px_3px_6px_#000000]" 
            onClick={onClick}
        >
            {children}
        </button>
    )
}

export default SidebarCircleButton