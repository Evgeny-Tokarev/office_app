import Axios from 'axios'
import {type RestApiResponse, type ErrorResponse} from '@/app/models'

export default function restApiErrorHandler<T>(err: unknown): RestApiResponse<T> {
    if (Axios.isAxiosError(err)) {
        const response = err.response;
        if (response) {
            console.log("error response:", response)
            const errorResponse: ErrorResponse = {};
            if (response.data) {
                errorResponse.code = response.statusText;
                errorResponse.message = response.data.message;
            }
            return { error: errorResponse, data: undefined };
        }
    }
    return { error: { message: 'Unknown error' }, data: undefined };
}
