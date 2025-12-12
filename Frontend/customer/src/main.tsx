import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, Outlet, RouterProvider } from 'react-router';

import './index.css'
import Layout from './pages/layout';
import Main from './pages/index';
import Products from './pages/products/index';
import Product from './pages/products/product/index';
import Reviews from './pages/products/product/reviews';
import Vendor from './pages/vendor/index';
import Profile from './pages/profile/index';
import Notifications from './pages/notifications/index';

const router = createBrowserRouter([
    {
        path: '/',
        element: <Layout/>,
        children: [
            { 
                index: true, 
                element: <Main/> 
            },
            { 
                path: 'products', 
                element: <Products/>, 
            },
            { 
                path: 'products/:productId', 
                element: <Outlet />,
                children: [
                    { 
                        index: true, 
                        element: <Product/> 
                    },
                    { 
                        path: 'reviews', 
                        element: <Reviews/> 
                    },
                ]
            },
            { 
                path: 'vendor/:id', 
                element: <Vendor/> 
            },
            { 
                path: 'profile', 
                element: <Profile/> 
            },
            { 
                path: 'notifications', 
                element: <Notifications/> 
            },
        ]
    }
])

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <RouterProvider router={router} />
    </StrictMode>,
);
