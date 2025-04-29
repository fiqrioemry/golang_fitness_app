import {
  FaGithub,
  FaLinkedin,
  FaTwitter,
  FaInstagram,
  FaEnvelope,
} from "react-icons/fa";

const Footer = () => {
  return (
    <footer className="bg-background border-t text-muted-foreground">
      <div className="container mx-auto px-4 py-10 grid gap-8 grid-cols-1 md:grid-cols-4 ">
        {/* Brand & Slogan */}
        <div>
          <h2 className="text-xl font-bold text-primary">Ecommerce</h2>
          <p className="text-sm mt-2">
            Discover the best deals and shop your favorite products with ease.
          </p>
        </div>

        {/* Navigation Links */}
        <div>
          <h3 className="text-base font-semibold mb-2">Quick Links</h3>
          <ul className="space-y-1 text-sm">
            <li>
              <a href="/" className="hover:text-primary">
                Home
              </a>
            </li>
            <li>
              <a href="/products" className="hover:text-primary">
                Products
              </a>
            </li>
            <li>
              <a href="/about" className="hover:text-primary">
                About Us
              </a>
            </li>
            <li>
              <a href="/contact" className="hover:text-primary">
                Contact
              </a>
            </li>
          </ul>
        </div>

        {/* Newsletter */}
        <div>
          <h3 className="text-base font-semibold mb-2">Newsletter</h3>
          <p className="text-sm mb-3">
            Get the latest updates and offers right into your inbox.
          </p>
          <form className="flex flex-col gap-2">
            <input
              type="email"
              placeholder="Enter your email"
              className="px-3 py-2 rounded border focus:outline-none focus:ring focus:border-primary text-sm"
            />
            <button
              type="submit"
              className="bg-primary text-white text-sm px-4 py-2 rounded hover:bg-primary/90 transition"
            >
              Subscribe
            </button>
          </form>
        </div>

        {/* Social Media */}
        <div>
          <h3 className="text-base font-semibold mb-2">Connect with Us</h3>
          <div className="flex space-x-4 mt-2">
            <a
              href="https://github.com/yourusername"
              target="_blank"
              rel="noopener noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaGithub />
            </a>
            <a
              href="https://linkedin.com/in/yourusername"
              target="_blank"
              rel="noopener noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaLinkedin />
            </a>
            <a
              href="https://twitter.com/yourusername"
              target="_blank"
              rel="noopener noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaTwitter />
            </a>
            <a
              href="https://instagram.com/yourusername"
              target="_blank"
              rel="noopener noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaInstagram />
            </a>
            <a
              href="mailto:info@example.com"
              className="text-xl hover:text-primary transition"
            >
              <FaEnvelope />
            </a>
          </div>
        </div>
      </div>

      <div className="text-center text-xs text-muted-foreground mt-6 py-4 border-t">
        &copy; {new Date().getFullYear()} Ecommerce. Built with ReactJS & Go.
      </div>
    </footer>
  );
};

export default Footer;
