function CircleButton({children, onClick, className}) {
    return (
        <button 
            className={`p-2.5 rounded-full flex items-center justify-center text-sidebar-fourth/60 transition duration-500 border border-sidebar-fourth/20 shadow-[3px_3px_6px_#000000] cursor-pointer hover:bg-sidebar-fourth/10 hover:text-sidebar-fourth hover:border-sidebar-fourth/40 ${className}`}
            onClick={onClick}
        >
            {children}
        </button>
    )
}

export default CircleButton