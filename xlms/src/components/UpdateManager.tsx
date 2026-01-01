import { Dialog, DialogPanel, DialogTitle } from "@headlessui/react";
import { useEffect, useMemo, useState } from "react";
import type { User } from "../types/user";
import api from "../api/api";
import toast from "react-hot-toast";

type DialogProps = {
    isOpen: boolean;
    user: User | undefined;
    managersList: User[];
    onClose: () => void;
    toast: any
};

export default function UpdateManager({ isOpen, onClose, user, managersList }: DialogProps) {
    const [search, setSearch] = useState("");
    const [selectedManager, setSelectedManager] = useState<User | null>(null);
    const [isOpenList, setIsOpenList] = useState(false);



    const saveChanges = async () => {
        try {
            const res = await api.post(`/managers/set?manager=${selectedManager?.id}&user=${user?.id}`)
            toast.success(res.data.message)
        } catch (err) {
            const message = err?.response?.data?.message || err?.response?.data?.error || err?.message || "Something went wrong";
            toast.error(message);
        }
        onClose();
    }

    const filteredManagers = useMemo(() => {
        const q = search.trim().toLowerCase();
        if (!q) return managersList;

        return managersList.filter(
            (m) =>
                m.full_name.toLowerCase().includes(q) ||
                m.email.toLowerCase().includes(q)
        );
    }, [search, managersList]);


    useEffect(() => {
        if (isOpen) {
            setSearch("");
            setSelectedManager(null);
            setIsOpenList(false);
        }
    }, [isOpen]);

    return (
        <Dialog open={isOpen} onClose={onClose} className="relative z-50">
            <div className="fixed inset-0 bg-black/30" />
            <div className="fixed inset-0 flex items-center justify-center p-4">
                <DialogPanel className="w-full max-w-md p-5 bg-white rounded-lg shadow-md">
                    <DialogTitle className="mb-3 text-sm font-semibold text-gray-800">
                        Update Manager for {user?.full_name}
                    </DialogTitle>

                    {/* Search */}
                    <input
                        type="text"
                        placeholder="Search manager..."
                        value={search}
                        onFocus={() => setIsOpenList(true)}
                        onChange={(e) => setSearch(e.target.value)}
                        className="w-full px-3 py-2 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-gray-400"
                    />

                    {/* Dropdown */}
                    {isOpenList && (
                        <div className="relative mt-1">
                            <div className="absolute z-10 w-full overflow-auto bg-white border border-gray-200 rounded-md shadow max-h-56">
                                {filteredManagers.length === 0 ? (
                                    <div className="px-3 py-2 text-sm text-gray-500">
                                        No managers found
                                    </div>
                                ) : (
                                    filteredManagers.map((manager) => (
                                        <div
                                            key={manager.id}
                                            onClick={() => {
                                                setSelectedManager(manager);
                                                setSearch(manager.full_name);
                                                setIsOpenList(false);
                                            }}
                                            className="relative px-3 py-2 text-sm cursor-pointer hover:bg-gray-100 group"
                                        >
                                            {manager.full_name}

                                            {/* Tooltip */}
                                            <span className="absolute z-20 hidden px-2 py-1 ml-2 text-xs text-white -translate-y-1/2 bg-gray-800 rounded left-full top-1/2 whitespace-nowrap group-hover:block">
                                                {manager.email}
                                            </span>
                                        </div>
                                    ))
                                )}
                            </div>
                        </div>
                    )}

                    {/* Actions */}
                    <div className="flex justify-end gap-2 mt-4">
                        <button
                            onClick={onClose}
                            className="px-3 py-1 text-sm border border-gray-300 rounded hover:bg-gray-100"
                        >
                            Cancel
                        </button>

                        <button
                            disabled={!selectedManager}
                            className="px-3 py-1 text-sm text-white bg-gray-800 rounded disabled:opacity-50"
                            onClick={saveChanges}
                        >
                            Save
                        </button>
                    </div>
                </DialogPanel>
            </div>
        </Dialog>
    );
}
