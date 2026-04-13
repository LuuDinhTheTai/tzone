import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { brandsApi } from '../../api/brands';
import { devicesApi } from '../../api/devices';
import { Tag, Smartphone, ArrowRight } from 'lucide-react';

export default function AdminDashboard() {
  const [brandCount, setBrandCount] = useState(0);
  const [deviceCount, setDeviceCount] = useState(0);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Promise.all([
      brandsApi.getAll(1, 1),
      devicesApi.getAll(1, 1),
    ]).then(([brandsRes, devicesRes]) => {
      setBrandCount(brandsRes.data.data?.pagination?.total || brandsRes.data.data?.total || 0);
      setDeviceCount(devicesRes.data.data?.pagination?.total || devicesRes.data.data?.total || 0);
    }).finally(() => setLoading(false));
  }, []);

  const cards = [
    {
      title: 'Total Brands',
      count: brandCount,
      icon: Tag,
      color: 'from-primary/20 to-primary/5',
      iconColor: 'text-primary',
      link: '/admin/brands',
    },
    {
      title: 'Total Devices',
      count: deviceCount,
      icon: Smartphone,
      color: 'from-accent/20 to-accent/5',
      iconColor: 'text-accent',
      link: '/admin/devices',
    },
  ];

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-text-primary">Dashboard</h1>
        <p className="text-sm text-text-muted mt-1">Overview of your TZone data</p>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 gap-5 mb-10">
        {cards.map((card) => {
          const Icon = card.icon;
          return (
            <div key={card.title} className="glass rounded-2xl p-6 card-hover">
              <div className="flex items-center justify-between mb-4">
                <div className={`w-11 h-11 rounded-xl bg-gradient-to-br ${card.color} flex items-center justify-center`}>
                  <Icon size={22} className={card.iconColor} />
                </div>
                <Link
                  to={card.link}
                  className="flex items-center gap-1 text-xs text-text-muted hover:text-primary transition-colors"
                >
                  Manage <ArrowRight size={14} />
                </Link>
              </div>
              <div>
                <p className="text-3xl font-bold text-text-primary">
                  {loading ? '—' : card.count}
                </p>
                <p className="text-sm text-text-muted mt-1">{card.title}</p>
              </div>
            </div>
          );
        })}
      </div>

      <div className="glass rounded-2xl p-6">
        <h2 className="text-lg font-semibold text-text-primary mb-4">Quick Actions</h2>
        <div className="flex flex-wrap gap-3">
          <Link
            to="/admin/brands"
            className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-medium text-text-primary border border-border hover:bg-surface-light transition-all"
          >
            <Tag size={16} /> Manage Brands
          </Link>
          <Link
            to="/admin/devices"
            className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-medium text-text-primary border border-border hover:bg-surface-light transition-all"
          >
            <Smartphone size={16} /> Manage Devices
          </Link>
        </div>
      </div>
    </div>
  );
}
