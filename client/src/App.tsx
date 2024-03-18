import CitizenList from "./components/CitizenList";
import Search from "./components/Search";
import { useEffect, useState } from "react";

function App() {

  const [searchTerm, setSearchTerm] = useState("");

  const citizens = [
    {
      "firstName": "John",
      "lastName": "Doe",
      "dateOfBirth": "1990-01-01",
      "gender": "Male",
      "address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "pincode": "10001"
    },
    {
      "firstName": "Jane",
      "lastName": "Doe",
      "dateOfBirth": "1990-01-01",
      "gender": "Female",
      "address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "pincode": "10001"
    },
    {
      "firstName": "Mary",
      "lastName": "Doe",
      "dateOfBirth": "1990-01-01",
      "gender": "Female",
      "address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "pincode": "10001"
    },
    {
      "firstName": "John",
      "lastName": "Doe",
      "dateOfBirth": "1990-01-01",
      "gender": "Male",
      "address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "pincode": "10001"
    },
    {
      "firstName": "Jane",
      "lastName": "Doe",
      "dateOfBirth": "1990-01-01",
      "gender": "Female",
      "address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "pincode": "10001"
    }
  ]


  useEffect(() => {

  }, [searchTerm]);

  const filteredCitizens = citizens.filter((citizen) => {
    return citizen.firstName.toLowerCase().includes(searchTerm.toLowerCase());
  });

  return (
    <div className="h-screen flex flex-col items-center">
      <Search searchTerm={searchTerm} setSearchTerm={setSearchTerm} />
      <CitizenList filteredCitizens={filteredCitizens} />
    </div>
  )
}

export default App
