import React from 'react'
import { IoBagHandle } from "react-icons/io5";
import { IoMdPeople } from "react-icons/io";
import { TiMessages } from "react-icons/ti";
import { IoSettingsOutline } from "react-icons/io5";

function BottomNav() {
    return (


        <div className="fixed z-50 w-full h-16 max-w-lg -translate-x-1/2 bg-white border border-gray-200 rounded-t-md bottom-0 left-1/2 dark:bg-gray-700 dark:border-gray-600">
            <div className="grid h-full max-w-lg grid-cols-4 mx-auto">
                <button type="button" className="inline-flex flex-col items-center justify-center px-5 rounded-s-full hover:bg-gray-50 dark:hover:bg-gray-800 group hover:blue-600 ">
                    <IoBagHandle className="group-hover:text-blue-600 text-stone-600" size={20} />
                    <p className="group-hover:text-blue-600 text-stone-600 text-xs">Home</p>
                </button>

                <button type="button" className="inline-flex flex-col items-center justify-center px-5 hover:bg-gray-50 dark:hover:bg-gray-800 group">
                    <IoMdPeople size={20} className="group-hover:text-blue-600 text-stone-600" />

                    <p className="group-hover:text-blue-600 text-stone-600 text-xs">Find Job</p>
                </button>



                <button type="button" className="inline-flex flex-col items-center justify-center px-5 hover:bg-gray-50 dark:hover:bg-gray-800 group">
                    <TiMessages size={20} className="group-hover:text-blue-600 text-stone-600" />
                    <p className="group-hover:text-blue-600 text-stone-600 text-xs">Mentorship</p>
                </button>

                <button type="button" className="inline-flex flex-col items-center justify-center px-5 rounded-e-full hover:bg-gray-50 dark:hover:bg-gray-800 group">
                    <IoSettingsOutline size={20} className="group-hover:text-blue-600 text-stone-600" />
                    <p className="group-hover:text-blue-600 text-stone-600 text-xs">Profile</p>
                </button>

            </div>
        </div>

    )
}

export default BottomNav