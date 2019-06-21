// in src/App.js
import React from 'react';
import { Admin, Resource } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import { CompetitionResultsList } from './js/components/container/competition_results';
import dataProvider from './js/dataProvider';

// const dataProvider = jsonServerProvider('http://jsonplaceholder.typicode.com');
const App = () => (
    <Admin dataProvider={dataProvider}>
    <Resource name="competition_results" list={CompetitionResultsList} />
    </Admin>
)

export default App;