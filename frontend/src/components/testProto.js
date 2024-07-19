import React from 'react';
import * as proto from '../utils/gamestate_pb';

const TestComponent = () => {
    console.log('proto:', proto);
    console.log('proto.gamestate:', proto.gamestate);
    return <div>Check console for proto object</div>;
};

export default TestComponent;
