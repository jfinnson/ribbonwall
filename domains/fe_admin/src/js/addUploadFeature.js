// addUploadFeature.js

const convertFileToBase64 = file => new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file.rawFile);

    reader.onload = () => resolve(reader.result);
    reader.onerror = reject;
});

const addUploadFeature = requestHandler => (type, resource, params) => {

    if (type === 'CREATE' && resource === 'competition_results/upload') {

        if (params.data.competition_results) {

            // NEW CODE HERE (to upload just one file):
            const myFile = params.data.competition_results;
            if ( !myFile.rawFile instanceof File ){
                return Promise.reject('Error: Not a file...'); // Didn't test this...
            }

            return Promise.resolve( convertFileToBase64(myFile) )
                .then( (picture64) => ({
                    src: picture64,
                    title: `${myFile.title}`
                }))
                .then( transformedMyFile => requestHandler(type, resource, {
                    ...params,
                    data: {
                        ...params.data,
                        competition_results: transformedMyFile
                    }
                }));
        }
    }
    return requestHandler(type, resource, params);
};

export default addUploadFeature;