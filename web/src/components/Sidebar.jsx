import uploadImg from "../assets/upload.png"
import home from "../assets/house-regular-full.svg"

function Sidebar() {
    return(
        <div className="w-64 h-full bg-sidebar-primary rounded-4xl border border-sidebar-fourth/20">
            <div className="flex items-center gap-4 py-8 px-5">
                <img src={uploadImg} alt="logo" className="w-8 h-8" />
                <p className="text-sidebar-trinary font-medium">My Project</p>
            </div>
            <div className="flex flex-col gap-y-1 p-3">
                <button className="w-full h-10 bg-sidebar-trinary rounded-xl cursor-pointer text-left px-3.5 flex items-center gap-2.5">
                    <img src={home} alt="home" className="w-5 h-5" />
                    <p className="text-sm font-medium">Home</p>
                </button>
            </div>
        </div>
    )
}

export default Sidebar