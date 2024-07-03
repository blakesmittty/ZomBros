import React, { useState } from "react"

function CharacterSelect({ onSelectCharacter }) {
    const characters = ['Character1', 'Character2', 'Character3', 'Character4'];

    return (
        <div>
            <h3>Select Your Character</h3>
            {characters.map((character) => (
                <button key={character} onClick={() => onSelectCharacter(character)}>
                    {character}
                </button>
            ))}
        </div>
    );
};

export default CharacterSelect;