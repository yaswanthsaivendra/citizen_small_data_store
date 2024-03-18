

const headings = ["First name", "Last name", "Date of birth", "Gender", "Address", "City", "State", "Pincode", "Edit", "Delete"]


const CitizenList = ({ filteredCitizens }) => {

    return (
        <div className="mt-12">
            <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
                <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                    <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            {headings.map((heading, index) => (
                                <th key={index} scope="col" className="px-6 py-3">
                                    {heading}
                                </th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {filteredCitizens.map((citizen, index) => (
                            <tr key={index} className="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700">
                                <th scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    {citizen.firstName}
                                </th>
                                <td className="px-6 py-4">
                                    {citizen.lastName} 
                                </td>
                                <td className="px-6 py-4">
                                    {citizen.dateOfBirth}
                                </td>
                                <td className="px-6 py-4">
                                    {citizen.gender}
                                </td>
                                <td className="px-6 py-4">
                                    {citizen.address}
                                </td>
                                <td className="px-6 py-4">
                                    {citizen.city}
                                </td>
                                <td className="px-6 py-4">
                                    {citizen.state}
                                </td>
                                <td className="px-6 py-4">
                                    {citizen.pincode}
                                </td>
                                <td className="px-6 py-4">
                                    <a href="#" className="font-medium text-blue-600 dark:text-blue-500 hover:underline">Edit</a>
                                </td>
                                <td className="px-6 py-4">
                                    <a href="#" className="font-medium text-blue-600 dark:text-blue-500 hover:underline">Delete</a>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    )
}

export default CitizenList