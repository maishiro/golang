import {useState,useEffect} from 'react';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {EventsOn} from "../wailsjs/runtime/runtime";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    const [itemList, setItemList] = useState<string[]>([]);

    function greet() {
        Greet(name).then(updateResultText);
    }

    EventsOn( "message", ( msg: string ) => {
        console.log( "message", msg );

        let list = [];
        list.push( msg );
        itemList.forEach( item => {
            list.push( item );
        });

        setItemList( list );
    } );

    return (
        <div id="App">
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={greet}>Greet</button>
            </div>

            <ul>
                { itemList.map( (item: string) => (
                    <li key={item}>{item}</li>
                ))}
            </ul>
        </div>
    )
}

export default App
