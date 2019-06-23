import React, { Component } from 'react';
import { fetchUtils, Admin, Resource } from 'react-admin';
import CompetitionResults from './components/container';
import dataProvider from './dataProvider';
import addUploadCapabilities from './addUploadFeature';
import { config } from '../constants';

const API_URL = config.url.API_URL + '/api_admin/v1';

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }
    const token = localStorage.getItem('token');
    options.headers.set('Authorization', `Bearer ${token}`);
    return fetchUtils.fetchJson(url, options);
};

const uploadCapableClient = addUploadCapabilities(dataProvider(API_URL, httpClient));

class AdminBrowser extends Component {
    login() {
        this.props.auth.login();
    }
    render() {
        const { isAuthenticated } = this.props.auth;
        return (
            <div className="container">
                {
                    isAuthenticated() && (
                        <Admin dataProvider={uploadCapableClient}>
                            <Resource name="competition_results" list={CompetitionResults.list}/>
                        </Admin>
                    )
                }
                {
                    !isAuthenticated() && (
                        <h4>
                            You are not logged in! Please{' '}
                            <a style={{ cursor: 'pointer' }}
                               onClick={this.login.bind(this)}>
                                Log In
                            </a>
                            {' '}to continue.
                        </h4>
                    )
                }
            </div>
        );
    }
}

export default AdminBrowser;