import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, Outlet, RouterProvider } from 'react-router';

import './index.css'
import Main from './pages/Index';
import Products from './pages/products/Index';
import Product from './pages/products/product/Index';
import Reviews from './pages/products/product/Reviews';
import Vendor from './pages/vendor/Index';
import Profile from './pages/profile/Index';
import Notifications from './pages/notifications/Index';
import Layout from './pages/layout';

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
                    { index: true, element: <Product/> },
                    { path: 'reviews', element: <Reviews/> },
                ]
            },
            { 
                path: 'vendor', 
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
