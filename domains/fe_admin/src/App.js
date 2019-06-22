// in src/App.js
import React from 'react';
import { Admin, Resource } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import CompetitionResults from './js/components/container';
import dataProvider from './js/dataProvider';
// import addUploadCapabilities from './js/addUploadFeature';

// const dataProvider = jsonServerProvider('http://jsonplaceholder.typicode.com');


// const uploadCapableClient = addUploadCapabilities(dataProvider);

const App = () => (
    <Admin dataProvider={dataProvider}>
    <Resource name="competition_results" list={CompetitionResults.list}/>
    </Admin>
)

export default App;