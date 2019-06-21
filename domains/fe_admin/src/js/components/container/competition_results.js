// in src/users.js
import React from 'react';
import { List, Datagrid, TextField, EmailField } from 'react-admin';

export const CompetitionResultsList = props => (
    <List {...props}>
        <Datagrid rowClick="edit">
            <TextField source="uuid" />
            <TextField source="placing" />
        </Datagrid>
    </List>
);