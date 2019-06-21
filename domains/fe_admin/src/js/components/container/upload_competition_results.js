import React from 'react';
import { Create, TextInput, FileInput, FileField, SimpleForm, required } from 'react-admin';

class UploadCompetitionResults extends React.Component {
    save() {
        console.log("Test")
    }
    render() {
        const { push, classes, ...props } = this.props;
        return (
            <SimpleForm save={this.save}>
                <TextInput source="organization" validate={required()} />
                <FileInput source="competition_results" label="Competition Result CSV" validate={required()}>
                    <FileField source="src" title="title" />
                </FileInput>
            </SimpleForm>
        );
    }
}
export default UploadCompetitionResults;