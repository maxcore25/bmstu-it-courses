import { BranchesSection } from './ui/branches-section';
import { CoursesSection } from './ui/courses-section';
import { HeroSection } from './ui/hero-section';
import { PublicHeader } from './ui/public-header';
import { TutorsSection } from './ui/tutors-section';

export const LandingPage = () => {
  return (
    <>
      <PublicHeader />
      <main className='container mx-auto'>
        <HeroSection />
        <CoursesSection />
        <TutorsSection />
        <BranchesSection />
      </main>
    </>
  );
};
