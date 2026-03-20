import { Plus, Minus, ChevronLeft, ChevronRight, Package} from 'lucide-react'

function Home() {
    return (
        <div className="w-full h-full flex flex-col gap-3">
            <div className="gap-6 flex flex-col">
                <div>
                    <p className="text-4xl text-sidebar-trinary font-bold">Hi, User!</p>
                    <p className="text-sidebar-fourth/60 text-sm font-medium">Let's do something today</p>
                </div>
                <div className="grid grid-cols-[400px_repeat(2,1fr)] grid-rows-[auto_auto_auto] gap-3">
                    <div className="bg-sidebar-primary border border-sidebar-fourth/20 backdrop-blur-2xl backdrop-saturate-200 backdrop-brightness-110 h-auto rounded-2xl p-4 flex flex-col justify-center gap-1 overflow-hidden">
                        <p className="text-sidebar-fourth/60 text-sm font-medium">Current balanced</p>
                        <p className="text-sidebar-fourth text-2xl">Rp. 24.020.060</p>
                        <button className="mt-6 h-10 rounded-2xl w-full text-sidebar-fourth border border-sidebar-fourth/20 bg-linear-to-tl from-sidebar-trinary/20 to-sidebar-primary backdrop-blur-2xl cursor-pointer text-sm font-medium">Wallet Detail</button>
                    </div>
                    <div className="bg-sidebar-primary h-auto rounded-2xl row-span-3 col-span-2 border border-sidebar-fourth/20 overflow-hidden backdrop-blur-2xl backdrop-saturate-200 backdrop-brightness-110"></div>
                    <div className="bg-sidebar-primary border border-sidebar-fourth/20 backdrop-blur-2xl backdrop-saturate-200 backdrop-brightness-110 rounded-2xl p-4 flex flex-col justify-center overflow-hidden row-span-2">
                        <div className="gap-y-6 flex flex-col">
                            <div className="gap-1 mb-1">
                                <p className="text-sidebar-fourth/60 text-sm font-medium">Income balanced</p>
                                <div className='flex'>
                                    <div className="text-2xl text-sidebar-fourth flex items-center justify-center gap-4">
                                        <Plus size={15} strokeWidth={2.5} />
                                        <p>Rp. 24.020.060</p>
                                    </div>
                                </div>
                            </div>
                            <div className="gap-1 mb-1">
                                <p className="text-sidebar-fourth/60 text-sm font-medium">Outcome balanced</p>
                                <div className='flex'>
                                    <div className="text-2xl text-sidebar-fourth flex items-center justify-center gap-4">
                                        <Minus size={15} strokeWidth={2.5} />
                                        <p>Rp. 24.020.060</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <button className="mt-6 h-10 rounded-2xl w-full text-sidebar-fourth border border-sidebar-fourth/20 bg-linear-to-bl from-sidebar-trinary/20 to-sidebar-primary backdrop-blur-2xl cursor-pointer text-sm font-medium">Adding income</button>
                    </div>
                </div>
            </div>
            <div className='my-5 mx-auto flex gap-6 text-sidebar-fourth'>
                <button className='text-sidebar-fourth/60 cursor-pointer'>
                    <ChevronLeft size={22} strokeWidth={2.5} />
                </button>
                <p className='text-sidebar-fourth/60'>Wallet Dashboard</p>
                <button className='text-sidebar-fourth/60 cursor-pointer'>
                    <ChevronRight size={22} strokeWidth={2.5} />
                </button>
            </div>
            <div className="bg-sidebar-primary border border-sidebar-fourth/20 rounded-2xl overflow-hidden">
                <table className="w-full">
                    <thead>
                        <tr className="border-b border-sidebar-fourth/20">
                            <th className="text-left text-sidebar-fourth/60 text-sm font-medium p-4">Date</th>
                            <th className="text-left text-sidebar-fourth/60 text-sm font-medium p-4">Description</th>
                            <th className="text-left text-sidebar-fourth/60 text-sm font-medium p-4">Category</th>
                            <th className="text-right text-sidebar-fourth/60 text-sm font-medium p-4">Amount</th>
                            <th className="text-right text-sidebar-fourth/60 text-sm font-medium p-4">Status</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td colSpan={5} className="text-center text-sidebar-fourth/20 py-12">
                                <div className="flex flex-col items-center gap-3">
                                    <Package size={70} strokeWidth={1.2} className="text-sidebar-fourth/60" />
                                    <p className="text-sidebar-fourth/40 text-sm font-medium">No data found</p>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    )
}

export default Home