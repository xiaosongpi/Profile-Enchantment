import { ArrowRight, ArrowBigUpDash } from "lucide-react"

function Login() {
    return (
        <div className='w-full h-screen bg-clr-black grid grid-cols-[1fr_auto]'>
            <div className="bg-clr-black flex flex-col items-center justify-center">
                <div className="w-full flex flex-col items-center justify-center gap-10">
                    <div className="h-20 w-20 rounded-3xl bg-clr-gray flex items-center justify-center shadow-[3px_3px_6px_#000000]">
                        <ArrowBigUpDash className="h-10 w-10 text-clr-white"></ArrowBigUpDash>
                    </div>
                    <div className="flex flex-col items-center justify-center gap-3">
                        <p className="text-clr-white text-2xl font-bold">Enchance Your Self</p>
                        <p className="text-clr-white/60">Keep up your discipline</p>
                    </div>
                    <div className="px-10 w-full h-20 flex items-center justify-center gap-6">
                        <div className="max-w-96 w-full h-full bg-clr-gray rounded-xl flex items-center justify-center text-sm text-clr-white/60 shadow-[3px_3px_6px_#000000]">Free to get started</div>
                        <div className="max-w-96 w-full h-full bg-clr-gray rounded-xl flex items-center justify-center text-sm text-clr-white/60 shadow-[3px_3px_6px_#000000]">less than 100ms response time</div>
                        <div className="max-w-96 w-full h-full bg-clr-gray rounded-xl flex items-center justify-center text-sm text-clr-white/60 shadow-[3px_3px_6px_#000000]">4.9/5 star rating</div>
                    </div>
                </div>
            </div>
            <div className="min-w-md bg-clr-black border-l border-l-clr-white/10 flex flex-col items-center justify-center z-10">
                <div className="w-full px-10 gap-9 flex flex-col">
                    <div className="flex flex-col gap-1">
                        <p className="text-clr-white text-2xl font-bold">Welcome User</p>
                        <p className="text-sm flex gap-1 text-clr-white/10">Don't have account?<a className="flex items-center justify-center cursor-pointer text-clr-white/60">Sign up <ArrowRight className="w-3 h-3"></ArrowRight> </a></p>
                    </div>
                    <form action="" className="w-full gap-6 flex flex-col">
                        <div className="flex flex-col gap-1">
                            <label htmlFor="email" className="text-sm text-clr-white/60">EMAIL ADDRESS</label>
                            <input type="text" placeholder="Input email here" className="w-full h-12 px-3 border text-clr-white border-clr-white/10 bg-clr-gray rounded-xl shadow-[inset_3px_3px_6px_#000000] outline-none placeholder:text-clr-white/20" />
                        </div>
                        <div className="flex flex-col gap-1">
                            <label htmlFor="password" className="text-sm text-clr-white/60">PASSWORD</label>
                            <input type="text" placeholder="Input password here" className="w-full h-12 px-3 border text-clr-white border-clr-white/10 bg-clr-gray rounded-xl shadow-[inset_3px_3px_6px_#000000] outline-none placeholder:text-clr-white/20" />
                        </div>
                        <div className="flex flex-col gap-2">
                            <div className="flex gap-2 items-center">
                                <input type="radio" className="h-3 w-3 border border-clr-white/10 rounded-sm cursor-pointer" />
                                <p className="text-sm text-clr-white/60">Remember me</p>
                            </div>
                            <button className="h-12 bg-clr-white rounded-xl text-sm shadow-[3px_3px_6px_#000000] cursor-pointer">Login</button>
                        </div>
                    </form>
                </div>
                <p className="mt-3 text-clr-white/20 text-sm">Forgot password? <a className="text-clr-white/60 cursor-pointer">Click here</a></p>
            </div>
        </div>
    )
}

export default Login