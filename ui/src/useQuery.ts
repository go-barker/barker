import qs from 'querystring';
import { useLocation } from 'react-router-dom';

export function useQuery() {
    const search = useLocation().search.replace(/^\?/, '');
    return qs.parse(search);
}
