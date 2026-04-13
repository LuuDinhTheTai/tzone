import { Smartphone } from 'lucide-react';
import { Link } from 'react-router-dom';

export default function Footer() {
  return (
    <footer className="border-t border-border bg-surface/80">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          {/* Brand */}
          <div className="md:col-span-1">
            <div className="flex items-center gap-2 mb-4">
              <div className="w-8 h-8 rounded-lg btn-gradient flex items-center justify-center">
                <Smartphone size={18} className="text-white" />
              </div>
              <span className="text-lg font-bold gradient-text">TZone</span>
            </div>
            <p className="text-sm text-text-muted leading-relaxed">
              Your ultimate platform for comparing smartphone specifications. Make informed decisions with detailed side-by-side comparisons.
            </p>
          </div>

          {/* Quick Links */}
          <div>
            <h4 className="text-sm font-semibold text-text-primary mb-4">Quick Links</h4>
            <ul className="space-y-2.5">
              {[
                { to: '/', label: 'Home' },
                { to: '/brands', label: 'All Brands' },
                { to: '/compare', label: 'Compare' },
              ].map((link) => (
                <li key={link.to}>
                  <Link
                    to={link.to}
                    className="text-sm text-text-muted hover:text-primary transition-colors"
                  >
                    {link.label}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          {/* Account */}
          <div>
            <h4 className="text-sm font-semibold text-text-primary mb-4">Account</h4>
            <ul className="space-y-2.5">
              {[
                { to: '/login', label: 'Login' },
                { to: '/register', label: 'Register' },
              ].map((link) => (
                <li key={link.to}>
                  <Link
                    to={link.to}
                    className="text-sm text-text-muted hover:text-primary transition-colors"
                  >
                    {link.label}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          {/* Contact */}
          <div>
            <h4 className="text-sm font-semibold text-text-primary mb-4">About</h4>
            <p className="text-sm text-text-muted leading-relaxed">
              Built with Go, React & TypeScript. 
              Powered by MongoDB & Supabase.
            </p>
          </div>
        </div>

        <div className="border-t border-border mt-10 pt-6 text-center">
          <p className="text-xs text-text-muted">
            © {new Date().getFullYear()} TZone. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  );
}
